using System.Collections.Generic;
using System.Linq;

namespace Web.Wizards
{
	public class WizardSteps
	{
		private readonly List<WizardStep> _stepsList;

		private WizardStep _head;

		public WizardSteps(IEnumerable<WizardStep> steps, int active = 0)
		{
			_stepsList = new List<WizardStep>();
			steps.ToList().ForEach(Add);
			var activeStep = _stepsList[active];
			for (
				var node = activeStep;
				node?.PreviousStep != null;
				node = node.PreviousStep
			)
			{
				node.PreviousStep.Passed = true;
			}
			activeStep.Active = true;
		}

		public WizardStep ActiveStep => _stepsList.Last(step => step.Active);
	

		private void Add(WizardStep step)
		{
			_stepsList.Add(step);
			if (_head == null)
			{
				_head = step;
				return;
			}
			_head.NextStep = step;
			step.PreviousStep = _head;
			_head = step;
		}

		public IEnumerator<WizardStep> GetEnumerator()
		{
			for (var step = _stepsList.First(); step != null; step = step.NextStep)
			{
				yield return step;
			}
		}
	}
}